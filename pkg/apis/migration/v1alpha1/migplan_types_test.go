/*
Copyright 2019 Red Hat Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"fmt"
	"testing"

	"github.com/onsi/gomega"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestStorageMigPlan(t *testing.T) {
	key := types.NamespacedName{
		Name:      "foo",
		Namespace: "default",
	}
	created := &MigPlan{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: "default",
		}}
	g := gomega.NewGomegaWithT(t)

	// Test Create
	fetched := &MigPlan{}
	g.Expect(c.Create(context.TODO(), created)).NotTo(gomega.HaveOccurred())

	g.Expect(c.Get(context.TODO(), key, fetched)).NotTo(gomega.HaveOccurred())
	g.Expect(fetched).To(gomega.Equal(created))

	// Test Updating the Labels
	updated := fetched.DeepCopy()
	updated.Labels = map[string]string{"hello": "world"}
	g.Expect(c.Update(context.TODO(), updated)).NotTo(gomega.HaveOccurred())

	g.Expect(c.Get(context.TODO(), key, fetched)).NotTo(gomega.HaveOccurred())
	g.Expect(fetched).To(gomega.Equal(updated))

	// Test Delete
	g.Expect(c.Delete(context.TODO(), fetched)).NotTo(gomega.HaveOccurred())
	g.Expect(c.Get(context.TODO(), key, fetched)).To(gomega.HaveOccurred())
}

func TestPersistentVolume_Update(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	// Setup
	pvA := PersistentVolume{
		Name:         "Elvis",
		StorageClass: "local",
		Action:       "Copy",
	}
	pvB := PersistentVolume{
		Name:         "Elvis",
		StorageClass: "changed",
		Action:       "Copy",
	}

	// Test
	pvA.Update(pvB)

	// Validation
	g.Expect(pvA.staged).To(gomega.BeTrue())
	g.Expect(pvA.StorageClass).To(gomega.Equal(pvB.StorageClass))
}

func TestPersistentVolumes_AddPv(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	pvs := PersistentVolumes{}
	// Add First
	pvA := PersistentVolume{
		Name:         "Elvis",
		StorageClass: "local",
	}
	// Add pvA
	pvs.AddPv(pvA)
	// Validate pvA
	_, found := pvs.index[pvA.Name]
	g.Expect(found).To(gomega.BeTrue())
	g.Expect(pvs.List[0]).To(gomega.Equal(PersistentVolume{
		Name:         "Elvis",
		StorageClass: "local",
		staged:       true,
	}))

	// Add Second.
	pvB := PersistentVolume{
		Name:         "Another",
		StorageClass: "mounted",
	}
	// Add pvB
	pvs.AddPv(pvB)
	// Validate pvA & pvB
	_, found = pvs.index[pvA.Name]
	g.Expect(found).To(gomega.BeTrue())
	_, found = pvs.index[pvB.Name]
	g.Expect(found).To(gomega.BeTrue())
	g.Expect(pvs.List[0]).To(gomega.Equal(PersistentVolume{
		Name:         "Elvis",
		StorageClass: "local",
		staged:       true,
	}))
	g.Expect(pvs.List[1]).To(gomega.Equal(PersistentVolume{
		Name:         "Another",
		StorageClass: "mounted",
		staged:       true,
	}))
}

func TestPersistentVolumes_DeletePv(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	// Setup
	n := 5
	pvs := PersistentVolumes{}
	for i := 0; i < n; i++ {
		pv := PersistentVolume{
			Name:         fmt.Sprintf("%d", i),
			StorageClass: "local",
			Action:       "Copy",
		}
		pvs.AddPv(pv)
	}
	g.Expect(len(pvs.List)).To(gomega.Equal(n))

	// Test
	pvs.DeletePv("1", "3")

	// Validation
	g.Expect(pvs.List).To(gomega.Equal(
		[]PersistentVolume{
			{
				Name:         "0",
				StorageClass: "local",
				Action:       "Copy",
				staged:       true,
			},
			{
				Name:         "2",
				StorageClass: "local",
				Action:       "Copy",
				staged:       true,
			},
			{
				Name:         "4",
				StorageClass: "local",
				Action:       "Copy",
				staged:       true,
			},
		}))
}

func TestPersistentVolumes_FindPv(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	// Setup
	pvs := PersistentVolumes{}
	pv := PersistentVolume{
		Name:         "Elvis",
		StorageClass: "local",
		staged:       true,
	}
	pvs.AddPv(pv)

	// Test
	found := pvs.FindPv(pv)

	// Validation
	g.Expect(found).NotTo(gomega.Equal(nil))
	g.Expect(*found).To(gomega.Equal(pv))
}

func TestPersistentVolumes_BeginPvStaging(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	// Setup
	n := 5
	pvs := PersistentVolumes{}
	for i := 0; i < n; i++ {
		pv := PersistentVolume{
			Name:         fmt.Sprintf("%d", i),
			StorageClass: "local",
			Action:       "Copy",
		}
		pvs.AddPv(pv)
	}
	g.Expect(pvs.staging).To(gomega.BeFalse())
	for i := 0; i < n; i++ {
		g.Expect(pvs.List[i].staged).To(gomega.BeTrue())
	}

	// Test
	pvs.BeginPvStaging()

	// Validation
	g.Expect(pvs.staging).To(gomega.BeTrue())
	for i := 0; i < n; i++ {
		g.Expect(pvs.List[i].staged).To(gomega.BeFalse())
	}

}

func TestPersistentVolumes_EndPvStaging(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	// Setup
	pvs := PersistentVolumes{}
	for i := 0; i < 10; i++ {
		pv := PersistentVolume{
			Name:         fmt.Sprintf("%d", i),
			StorageClass: "local",
			Action:       "Copy",
		}
		pvs.AddPv(pv)
	}
	pvs.BeginPvStaging()

	// Test
	pvs.List[2].staged = true
	pvs.List[4].staged = true
	pvs.EndPvStaging()

	// Validation
	g.Expect(pvs.List).To(gomega.Equal(
		[]PersistentVolume{
			{
				Name:         "2",
				StorageClass: "local",
				Action:       "Copy",
				staged:       true,
			},
			{
				Name:         "4",
				StorageClass: "local",
				Action:       "Copy",
				staged:       true,
			},
		}))
}
