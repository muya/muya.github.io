Fred Muya's github blog



## Running this Project
### Without Docker
1. Set Up Jekyll as outlined here: https://jekyllrb.com/docs/installation/#guides
2. `cd` into the project directory
3. Run `bundle install` to install missing gems.
4. Run the command below:
```shell
bundle exec jekyll serve
```

### Using Docker
1. Ensure you have Docker installed (https://www.docker.com/get-started)
2. `cd` into the project directory
3. Run docker-compose up:
```shell
docker-compose -f deploy/docker-compose.yml up
```
4. Access the project on http://127.0.0.1:4000
