package app

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/AyodejiO/okteto/service"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func sanityCheck() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
}

func getK8sClient() *kubernetes.Clientset {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		panic(err.Error())
	}

	return clientset
}

func writeResponse(w http.ResponseWriter, code int, success bool, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	output := map[string]interface{}{
		"success": success,
		"data":    data,
	}
	if err := json.NewEncoder(w).Encode(output); err != nil {
		panic(err)
	}
}

func GetIndexHandler(w http.ResponseWriter, r *http.Request) {
	writeResponse(
		w, 
		http.StatusOK, 
		true, 
		map[string]interface{}{
			"message": "Welcome to the Okteto API",
			"timestamp": time.Now().Format(time.RFC3339),
		},
	)
}

func Custom404Handler(w http.ResponseWriter, r *http.Request) {
	writeResponse(
		w, 
		http.StatusNotFound, 
		false, 
		map[string]interface{}{
			"message": "Requested resource not found",
		},
	)
}

func Start() {
	sanityCheck()
	clientset := getK8sClient()

	ph := PodHandler{service.NewPodService(clientset)}

	router := mux.NewRouter()

	router.HandleFunc("/", GetIndexHandler).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(Custom404Handler)
	router.HandleFunc("/pods/count", ph.GetPodsCount).Methods("GET")
	router.HandleFunc("/pods", ph.GetPods).Methods("GET")

	// start the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", router))
}