package file_test

import (
	"io/ioutil"
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"encoding/json"

	"path/filepath"

	. "github.com/solo-io/gloo-testing/helpers"
	"github.com/solo-io/gloo/pkg/artifactwatcher"
	. "github.com/solo-io/gloo/pkg/artifactwatcher/file"
	"github.com/solo-io/gloo/pkg/log"
)

var _ = Describe("FileArtifactWatcher", func() {
	var (
		dir   string
		file  string
		ref   string
		err   error
		watch artifactwatcher.Interface
	)
	BeforeEach(func() {
		ref = "artifacts.yml"
		dir, err = ioutil.TempDir("", "fileartifacttest")
		Must(err)
		file = filepath.Join(dir, ref)
		watch, err = NewArtifactWatcher(dir, time.Millisecond)
		Must(err)
	})
	AfterEach(func() {
		log.Debugf("removing " + dir)
		os.RemoveAll(dir)
	})
	Describe("watching file", func() {
		Context("no artifacts wanted", func() {
			It("doesnt send anything on any channel", func() {
				missingArtifacts := map[string]map[string][]byte{"another-key": {"foo": []byte("bar"), "baz": []byte("qux")}}
				data, err := json.Marshal(missingArtifacts)
				Expect(err).NotTo(HaveOccurred())
				err = ioutil.WriteFile(file, data, 0644)
				Expect(err).NotTo(HaveOccurred())
				select {
				case <-watch.Artifacts():
					Fail("Artifacts was received, expected timeout")
				case err := <-watch.Error():
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 1):
					// passed
				}
			})
		})
		Context("valid artifacts are written to the ref file", func() {
			It("sends a corresponding Artifacts on Artifacts()", func() {

				data := []byte("this is the data")
				err = ioutil.WriteFile(file, data, 0644)
				Must(err)
				go watch.TrackArtifacts([]string{ref})
				select {
				case parsedArtifacts := <-watch.Artifacts():
					Expect(parsedArtifacts).To(Equal(artifactwatcher.Artifacts{
						ref: map[string][]byte{ref: data},
					}))
				case err := <-watch.Error():
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 5):
					Fail("expected new artifacts to be read in before 1s")
				}
			})
		})
	})
})
