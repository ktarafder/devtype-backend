from fastapi import FastAPI, HTTPException
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel
import pandas as pd
import joblib
import matplotlib.pyplot as plt
import base64
from io import BytesIO
from recommendations import RECOMMENDATIONS
import random

app = FastAPI(strict_slashes=False)

model = joblib.load("random_forest_model.pkl")

# Define request schemas with Pydantic
class InputData(BaseModel):
    overall_speed: float
    overall_accuracy: float

class TypingSessionData(BaseModel):
    overall_speed: float
    overall_accuracy: float

@app.post("/predict")
def predict_cluster(data: list[InputData]):
    try:
        input_df = pd.DataFrame([{"WPM": obj.overall_speed / 5, "Accuracy": obj.overall_accuracy} for obj in data])
        predicted_clusters = model.predict(input_df)

        most_frequent_cluster = pd.Series(predicted_clusters).mode()[0]

        recommendations = RECOMMENDATIONS[most_frequent_cluster]
        improvement_area = recommendations["improvement_area"]
        feedback_text = random.choice(recommendations["feedback_text"])

        response = {
            "cluster": int(most_frequent_cluster),
            "improvement_area": improvement_area,
            "feedback_text": feedback_text
        }

        return response
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

@app.post("/performance")
async def generate_performance_graphs(data: list[TypingSessionData]):
    try:
        # Convert the input data into a pandas DataFrame
        df = pd.DataFrame([entry.dict() for entry in data])

        # Generate the accuracy graph
        accuracy_base64 = generate_graph(
            df['overall_accuracy'],
            title="Typing Accuracy Over Sessions",
            ylabel="Accuracy (%)",
            xlabel="Session Index"
        )

        # Generate the speed graph
        speed_base64 = generate_graph(
            df['overall_speed'],
            title="Typing Speed Over Sessions",
            ylabel="Speed (WPM)",
            xlabel="Session Index"
        )

        # Return the Base64-encoded graphs
        return {
            "accuracy_graph": accuracy_base64,
            "speed_graph": speed_base64
        }

    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error generating graphs: {str(e)}")

def generate_graph(data, title, ylabel, xlabel):
    """Generates a graph and returns its Base64-encoded representation."""
    plt.figure(figsize=(10, 6))
    plt.plot(data, marker="o", linestyle="-", color="blue")
    plt.title(title)
    plt.xlabel(xlabel)
    plt.ylabel(ylabel)
    plt.grid(True)

    # Save the plot to a BytesIO buffer
    buf = BytesIO()
    plt.savefig(buf, format="png")
    buf.seek(0)
    plt.close()

    # Encode the image in Base64
    return base64.b64encode(buf.getvalue()).decode("utf-8")

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)