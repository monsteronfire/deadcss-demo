from fastapi import FastAPI, Request
from agents.css_analyser import analyse_css

app = FastAPI()


@app.get("/")
def read_root():
    return {"Hello": "World"}


@app.get("/health")
def health_check():
    return {"status": "healthy"}


@app.post("/api/analyse-css")
async def analyse_css_endpoint(request: Request):
    results = await analyse_css({})

    return {
        "status": "success",
        "analysis": results
    }
