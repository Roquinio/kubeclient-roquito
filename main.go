package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	// Kubernetes Go Client packages
	"github.com/TwiN/go-color"              // Print Color
	"github.com/common-nighthawk/go-figure" // ASCII Art
	"github.com/k0kubun/pp/v3"              // Print Color V2
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// Global variables

// Functions
func main() {
	subCommand()

}

func subCommand() {
	// helpCmd := flag.NewFlagSet("help", flag.ExitOnError)
	argLenght := os.Args[1:]
	if len(argLenght) >= 1 {
		args := os.Args[1]
		switch args {
		case "help", "-h":
			// helpCmd.Parse(os.Args[2:])
			h := "helpppp"
			pp.Println(h)

		default:
			fmt.Printf(color.Red+"Unknown synthax '%s'\n"+color.Reset, args)
		}
	} else {
		emptyCMD()
	}
}

// Function activated when no args given
func emptyCMD() {
	fmt.Println("---------------------------------------------------")
	asciiArt := figure.NewFigure("Roquito", "", true)
	asciiArt.Print()
	fmt.Println("---------------------------------------------------")
	fmt.Println(color.Cyan + "Kubernetes client wrotes in Golang" + color.Reset)
	fmt.Println("---------------------------------------------------")

	var kubeconfig *string

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	nodes, _ := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})

	for _, node := range nodes.Items {
		fmt.Printf("%s\n", node.Name)
		for _, condition := range node.Status.Conditions {
			switch condition.Type {
			case "MemoryPressure":
				if condition.Status == "False" {
					fmt.Println(color.Green + "\tRAM: Healty" + color.Reset)
				} else if condition.Status == "True" {
					fmt.Println(color.Yellow + "\tRAM : Not Healty" + color.Reset)
				} else {
					fmt.Println(color.Red + "\tRAM : Unknown" + color.Reset)
				}

			case "DiskPressure":
				if condition.Status == "False" {
					fmt.Println(color.Green + "\tDisk : Healty" + color.Reset)
				} else if condition.Status == "True" {
					fmt.Println(color.Yellow + "\tDisk : Not Healty" + color.Reset)
				} else {
					fmt.Println(color.Red + "\tDisk : Unknown" + color.Reset)
				}

			case "PIDPressure":
				if condition.Status == "False" {
					fmt.Println(color.Green + "\tCPU : Healty" + color.Reset)
				} else if condition.Status == "True" {
					fmt.Println(color.Yellow + "\tCPU : Not Healty" + color.Reset)
				} else {
					fmt.Println(color.Red + "\tCPU : Unknown" + color.Reset)
				}
				// pp.Printf("\t%s: %s\n", condition.Type, condition.Status)
				// fmt.Print(node.Status.Conditions[3])
			case "Ready":
				if condition.Status == "True" {
					fmt.Println(color.Green + "\tNode Status : Ready" + color.Reset)
				} else if condition.Status == "False" {
					fmt.Println(color.Yellow + "\tNode Status : Not Ready" + color.Reset)
				} else {
					fmt.Println(color.Red + "\tNode Status : Unknown" + color.Reset)
				}

			}
		}
	}
}
