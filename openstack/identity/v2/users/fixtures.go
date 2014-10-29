package users

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/rackspace/gophercloud/testhelper"
	fake "github.com/rackspace/gophercloud/testhelper/client"
)

func MockListResponse(t *testing.T) {
	th.Mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "users":[
        {
            "id": "u1000",
						"name": "John Smith",
            "username": "jqsmith",
            "email": "john.smith@example.org",
            "enabled": true,
						"tenant_id": "12345"
        },
        {
            "id": "u1001",
						"name": "Jane Smith",
            "username": "jqsmith",
            "email": "jane.smith@example.org",
            "enabled": true,
						"tenant_id": "12345"
        }
    ]
}
  `)
	})
}

func MockCreateUser(t *testing.T) {
	th.Mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		th.TestJSONRequest(t, r, `
{
    "user": {
		    "name": "new_user",
		    "tenant_id": "12345",
				"enabled": false,
				"email": "new_user@foo.com"
    }
}
	`)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "user": {
        "name": "new_user",
        "tenant_id": "12345",
        "enabled": false,
        "email": "new_user@foo.com",
        "id": "c39e3de9be2d4c779f1dfd6abacc176d"
    }
}
`)
	})
}

func MockGetUser(t *testing.T) {
	th.Mux.HandleFunc("/users/new_user", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
		"user": {
				"name": "new_user",
				"tenant_id": "12345",
				"enabled": false,
				"email": "new_user@foo.com",
				"id": "c39e3de9be2d4c779f1dfd6abacc176d"
		}
}
`)
	})
}

func MockUpdateUser(t *testing.T) {
	th.Mux.HandleFunc("/users/c39e3de9be2d4c779f1dfd6abacc176d", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		th.TestJSONRequest(t, r, `
{
    "user": {
		    "name": "new_name",
		    "enabled": true,
		    "email": "new_email@foo.com"
    }
}
`)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
		"user": {
				"name": "new_name",
				"tenant_id": "12345",
				"enabled": true,
				"email": "new_email@foo.com",
				"id": "c39e3de9be2d4c779f1dfd6abacc176d"
		}
}
`)
	})
}

func MockDeleteUser(t *testing.T) {
	th.Mux.HandleFunc("/users/c39e3de9be2d4c779f1dfd6abacc176d", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})
}
