package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/tliron/puccini/ard"
	"github.com/tliron/puccini/clout"
	"github.com/tliron/puccini/tosca/compiler"
	"github.com/tliron/puccini/tosca/normal"
	"github.com/tliron/puccini/tosca/parser"
	"github.com/tliron/puccini/tosca/problems"
)

func TestParse(t *testing.T) {
	testParse(t, "tosca/artifacts.yaml", nil)
	testParse(t, "tosca/attributes.yaml", nil)
	testParse(t, "tosca/data-types.yaml", nil)
	testParse(t, "tosca/descriptions.yaml", nil)
	testParse(t, "tosca/dsl-definitions.yaml", nil)
	testParse(t, "tosca/functions.yaml", nil)
	testParse(t, "tosca/inputs-and-outputs.yaml", ard.Map{"ram": "1gib"})
	testParse(t, "tosca/interfaces.yaml", nil)
	testParse(t, "tosca/metadata.yaml", nil)
	testParse(t, "tosca/namespaces.yaml", nil)
	testParse(t, "tosca/policies-and-groups.yaml", nil)
	testParse(t, "tosca/requirements-and-capabilities.yaml", nil)
	testParse(t, "tosca/simple-for-nfv.yaml", nil)
	testParse(t, "tosca/source-and-target.yaml", nil)
	testParse(t, "tosca/substitution-mapping-client.yaml", nil)
	testParse(t, "tosca/substitution-mapping.yaml", nil)
	testParse(t, "tosca/unicode.yaml", nil)
	testParse(t, "tosca/workflows.yaml", nil)
	testParse(t, "javascript/constraints.yaml", nil)
	testParse(t, "javascript/exec.yaml", nil)
	testParse(t, "javascript/functions.yaml", nil)
	testParse(t, "kubernetes/bookinfo/bookinfo-simple.yaml", nil)
	testParse(t, "openstack/hello-world.yaml", nil)
	testParse(t, "bpmn/open-loop.yaml", nil)
	testParse(t, "cloudify/advanced-blueprint-example.yaml", ard.Map{
		"host_ip":                "1.2.3.4",
		"agent_user":             "my_user",
		"agent_private_key_path": "my_key",
	})
	testParse(t, "cloudify/example.yaml", nil)
	testParse(t, "hot/hello-world.yaml", ard.Map{
		"key_name":          "my_key",
		"image_id":          "my_image",
		"database_password": "A12345",
	})
	testParse(t, "hot/single-server-with-existing-floating-ip.yaml", ard.Map{
		"public_network": "public",
		"floating_ip":    "1.2.3.4",
		"image":          "my_image",
		"flavor":         "my_flavor",
		"ssh_keys":       "first,second",
	})
}

var ROOT string

func init() {
	ROOT = os.Getenv("ROOT")
}

func testParse(t *testing.T, url string, inputs ard.Map) {
	t.Run(url, func(t *testing.T) {
		// Running the tests in parallel is not for speed;
		// it actually allowed us to find several concurrency bugs
		t.Parallel()

		var s *normal.ServiceTemplate
		var c *clout.Clout
		var p *problems.Problems
		var err error

		if s, p, err = parser.Parse(fmt.Sprintf("%s/examples/%s", ROOT, url), nil, inputs); err != nil {
			t.Errorf("%s\n%s", err.Error(), p)
			return
		}

		if c, err = compiler.Compile(s); err != nil {
			t.Errorf("%s\n%s", err.Error(), p)
			return
		}

		compiler.Resolve(c, p, "yaml", true)
		if !p.Empty() {
			t.Errorf("%s", p)
			return
		}

		compiler.Coerce(c, p, "yaml", true)
		if !p.Empty() {
			t.Errorf("%s", p)
			return
		}
	})
}
