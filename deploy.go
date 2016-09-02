package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	client "k8s.io/kubernetes/pkg/client/unversioned"
)

func isRunning(kubeClient *client.Client, pod string, namespace string) {
	// get info
	// status := getPodStatus(kubeClient, pod, namespace)
	// fmt.Println(status)
	// check status

	// return
}

func replaceImage(pod, oldImage, newImage string) {

	commandOptions := []string{"get", "pod", pod, "-o", "yaml"}
	result := execOutput("kubectl", commandOptions)
	result = strings.Replace(result, oldImage, newImage, -1)

	ioutil.WriteFile("tmp.dat", []byte(result), os.ModePerm)
	defer os.Remove("tmp.dat")

	commandOptions = []string{"replace", "-f", "tmp.dat"}
	_, err := exec.Command("kubectl", commandOptions...).Output()
	if err != nil {
		log.Fatal(err)
	}
}

func deploy(kubeClient *client.Client, params map[string]string) {

	podInfos := getPodInfos(getPods(kubeClient, ""))

	myPodInfo := getMatchedPodInfo(params["pod"], podInfos)
	if myPodInfo["pod"] == "" {
		fmt.Println(params["pod"] + " doesn't exist.")
		os.Exit(1)
	}

	replaceImage(myPodInfo["pod"], myPodInfo["image"], params["image"])

	isRunning(kubeClient, params["pod"], params["namespace"])

}
