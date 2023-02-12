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

func GetIndex(w http.ResponseWriter, r *http.Request) {
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

func Start() {
	sanityCheck()
	clientset := getK8sClient()

	ph := PodHandlers{service.NewPodService(clientset)}

	router := mux.NewRouter()

	router.HandleFunc("/", GetIndex).Methods("GET")
	router.HandleFunc("/pods/count", ph.GetPodCount).Methods("GET")
	router.HandleFunc("/pods", ph.GetPods).Methods("GET")

	// start the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", router))
}