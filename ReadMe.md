## Mamas Kitchen:

### How to run locally:
In order to run Mama's Kitchen you will need to have docker and docker-compose installed. 

Run `docker compose -f 'docker-compose.yml' up -d --build `
> If you want to see logs you can omit the "-d"

Once all the containers are running you will need to install the models necessary.
1. Go to localhost:8080/browse.
2. Install phi-4 model.
3. Install whipser-1 model.
4. Navigate to localhost:8080/ to see the progress of the install.

Once the models are installed you can upload your recipe audio to the site: localhost:8000/


### Tasks

#### Backend
- [] Add Go Logging 
- [] Add Python Logging
- [] add error handling for python
- [] add error handling for golang
- [x] Create a InMemory Recipe Store to serve to the UI
- [x] local upload to blob store
- [x] local enqueue message
- [x] build worker(python) to pull message
- [x] transcribe locally
- [x] transform locally
- [] connect to a database(mongo)



#### Frontend
- [x] Create a Logo with a link to the grid view
- [x] upload component to upload audio files
- [] record audio and upload it to storage.
- [] style the page to make it look good
- [] search page


#### Clean up Tasks
- [x] restructure the code
- [] add linters and formatters to auto check during build
- [x] use docker to build frontend and a script to tie it in with go binary
- [x] add cmd folder
- [x] build scripts - docker compose
- [x] upload to github
- [] use cloud prod envs

