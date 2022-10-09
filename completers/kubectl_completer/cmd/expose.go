package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var exposeCmd = &cobra.Command{
	Use:   "expose",
	Short: "Take a replication controller, service, deployment or pod and expose it as a new Kubernetes service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(exposeCmd).Standalone()
	exposeCmd.Flags().Bool("allow-missing-template-keys", true, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	exposeCmd.Flags().String("cluster-ip", "", "ClusterIP to be assigned to the service. Leave empty to auto-allocate, or set to 'None' to create a headless service.")
	exposeCmd.Flags().String("dry-run", "none", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	exposeCmd.Flags().String("external-ip", "", "Additional external IP address (not managed by Kubernetes) to accept for the service. If this IP is routed to a node, the service can be accessed by this IP in addition to its generated service IP.")
	exposeCmd.Flags().String("field-manager", "kubectl-expose", "Name of the manager used to track field ownership.")
	exposeCmd.Flags().StringSliceP("filename", "f", []string{}, "Filename, directory, or URL to files identifying the resource to expose a service")
	exposeCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	exposeCmd.Flags().StringP("labels", "l", "", "Labels to apply to the service created by this call.")
	exposeCmd.Flags().String("load-balancer-ip", "", "IP to assign to the LoadBalancer. If empty, an ephemeral IP will be created and used (cloud-provider specific).")
	exposeCmd.Flags().String("name", "", "The name for the newly created object.")
	exposeCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	exposeCmd.Flags().String("override-type", "merge", "The method used to override the generated object: json, merge, or strategic.")
	exposeCmd.Flags().String("overrides", "", "An inline JSON override for the generated object. If this is non-empty, it is used to override the generated object. Requires that the object supply a valid apiVersion field.")
	exposeCmd.Flags().String("port", "", "The port that the service should serve on. Copied from the resource being exposed, if unspecified")
	exposeCmd.Flags().String("protocol", "", "The network protocol for the service to be created. Default is 'TCP'.")
	exposeCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	exposeCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.")
	exposeCmd.Flags().String("selector", "", "A label selector to use for this service. Only equality-based selector requirements are supported. If empty (the default) infer the selector from the replication controller or replica set.)")
	exposeCmd.Flags().String("session-affinity", "", "If non-empty, set the session affinity for the service to this; legal values: 'None', 'ClientIP'")
	exposeCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	exposeCmd.Flags().String("target-port", "", "Name or number for the port on the container that the service should direct traffic to. Optional.")
	exposeCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	exposeCmd.Flags().String("type", "", "Type for this service: ClusterIP, NodePort, LoadBalancer, or ExternalName. Default is 'ClusterIP'.")
	exposeCmd.Flag("dry-run").NoOptDefVal = "unchanged"
	rootCmd.AddCommand(exposeCmd)

	carapace.Gen(exposeCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":   action.ActionDryRunModes(),
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    action.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
		"type":      carapace.ActionValues("ClusterIP", "NodePort", "LoadBalancer", "ExternalName"),
	})

	carapace.Gen(exposeCmd).PositionalCompletion(
		carapace.ActionValues("pod", "service", "replicationcontroller", "deployment", "replicaset"),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionResources("", c.Args[0])
		}),
	)
}