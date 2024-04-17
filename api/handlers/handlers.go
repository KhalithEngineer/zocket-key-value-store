package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"zocket-key-value-store/internal/store"
	"zocket-key-value-store/internal/utils"

	"github.com/hashicorp/consul/api"
)

type PutRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// GetRequest represents the request payload for GET operation
type GetRequest struct {
	Key string `json:"key"`
}

// KeyValueHandler handles HTTP requests for key-value operations
type KeyValueHandler struct {
	store        *store.KeyValueStore
	consulClient *api.Client
}

func NewKeyValueHandler(consulClient *api.Client) *KeyValueHandler {
	return &KeyValueHandler{
		store:        store.NewKeyValueStore(),
		consulClient: consulClient,
	}
}

// PutHandler handles HTTP PUT requests
func (h *KeyValueHandler) PutHandler(w http.ResponseWriter, r *http.Request) {
	var req PutRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	kv := &api.KVPair{Key: req.Key, Value: []byte(req.Value)}

	if _, err := h.consulClient.KV().Put(kv, nil); err != nil {
		fmt.Println("Error replicating data to Consul:", err)
	}

	h.store.Put(req.Key, req.Value)

	w.WriteHeader(http.StatusCreated)
}

// GetHandler handles HTTP GET requests
func (h *KeyValueHandler) GetHandler(w http.ResponseWriter, r *http.Request) {

	var req GetRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	keyValue, _, err := h.consulClient.KV().Get(req.Key, &api.QueryOptions{})

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error from cluster")
		return
	}

	if keyValue == nil {
		utils.RespondWithError(w, http.StatusNotFound, "Key not found")
		return

	}

	response := map[string]string{"value": string(keyValue.Value)}
	json.NewEncoder(w).Encode(response)
}

// DeleteHandler handles HTTP DELETE requests
func (h *KeyValueHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	var req GetRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	h.store.Delete(req.Key)

	// Delete key from Consul KV store for replication
	if _, err := h.consulClient.KV().Delete(req.Key, nil); err != nil {
		fmt.Println("Error deleting key from Consul:", err)
	}

	w.WriteHeader(http.StatusNoContent)
}
