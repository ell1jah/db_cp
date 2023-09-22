from locust import HttpUser, task, between

def login(l):
    l.client.post("/login", json={"login":"crawfordjohn", "password":"qJ8OqIdW&j"})

class WebsiteTestUser(HttpUser):
    wait_time = between(0, 3.0)

    def on_start(self):
        login(self)

    @task(2)
    def hello_world(self):
        self.client.get(url="/basket", headers={"Authorization":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjo2MSwicm9sZSI6InVzZXIifSwiZXhwIjoxNjk1ODk5NTI1LCJpYXQiOjE2OTUyOTQ3MjV9.lyvBqUHnkp9sDxyrDXESL5X0d7uPwGBfwCiqaFxCAvw"})
