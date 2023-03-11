package api

// import (
// 	"database/sql"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	mockdb "github.com/nina/db/mock"
// 	db "github.com/nina/db/sqlc"
// 	"github.com/stretchr/testify/require"
// )

// func TestGetProject(t *testing.T) {
// 	testCases := []struct {
// 		name          string
// 		projectName   string
// 		buildStubs    func(store *mockdb.MockStore)
// 		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name:        "OK",
// 			projectName: "SupremeFinance",
// 			buildStubs: func(store *mockdb.MockStore) {
// 				ret := db.AddProjectParams{
// 					ProjectName: "SupremeFinance",
// 				}

// 				store.EXPECT().
// 					GetProject(gomock.Any(), gomock.Eq("SupremeFinance")).
// 					Times(1).
// 					Return(ret, nil)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusOK, recorder.Code)
// 			},
// 		},
// 		{
// 			name:        "Not found",
// 			projectName: "NotFound",
// 			buildStubs: func(store *mockdb.MockStore) {
// 				ret := db.AddProjectParams{}
// 				fmt.Println("test")
// 				store.EXPECT().
// 					GetProject(gomock.Any(), gomock.Eq("NotFound")).
// 					Times(1).
// 					Return(ret, sql.ErrNoRows)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusNotFound, recorder.Code)
// 			},
// 		},
// 	}

// 	for i := range testCases {
// 		test := testCases[i]
// 		t.Run(test.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			store := mockdb.NewMockStore(ctrl)
// 			test.buildStubs(store)

// 			server := NewServer(store)
// 			recorder := httptest.NewRecorder()

// 			url := fmt.Sprintf("/projects?project_name=%s", test.projectName)

// 			request, err := http.NewRequest(http.MethodGet, url, nil)

// 			require.NoError(t, err)

// 			server.router.ServeHTTP(recorder, request)
// 			fmt.Println(request)
// 		})

// 	}
// }

// // func TestAddProject(t *testing.T) {
// // 	ctrl := gomock.NewController(t)
// // 	defer ctrl.Finish()

// // 	addProject := db.CreateProjectInfoParams{
// // 		ProjectName: "Carmin Test",
// // 		Symbol:      "CARMIN",
// // 	}

// // 	store := mockdb.NewMockStore(ctrl)
// // 	store.EXPECT().
// // 		CreateProjectInfo(gomock.Any(), gomock.Eq(addProject)).
// // 		Times(1).
// // 		Return(addProject.ProjectName, nil)

// // 	server := NewServer(store)
// // 	recorder := httptest.NewRecorder()
// // 	encodedInput, err := json.Marshal(addProject)
// // 	if err != nil {
// // 		panic(err)
// // 	}
// // 	url := "/add-projects"
// // 	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(encodedInput))
// // 	if err != nil {
// // 		panic(err)
// // 	}
// // 	fmt.Println(request.Body)
// // 	require.NoError(t, err)

// // 	server.router.ServeHTTP(recorder, request)
// // 	require.Equal(t, http.StatusOK, recorder.Code)
// // }
