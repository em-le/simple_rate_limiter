from locust import HttpUser, task, between

class HelloWorldUser(HttpUser):
    root_url = "http://127.0.0.1:8080"
    wait_time = between(0.1, 0.3)

    @task
    def hello_world(self):
       response = self.client.get(f"{self.root_url}/health_check")
