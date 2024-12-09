from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
import pandas as pd
import joblib

app = FastAPI(strict_slashes=False)

model = joblib.load("random_forest_model.pkl")

# Define request schema with Pydantic
class InputData(BaseModel):
    WPM: float
    Accuracy: float

@app.post("/predict")
def predict_cluster(data: InputData):
    try:
        input_df = pd.DataFrame({
            "WPM": [data.WPM / 5],
            "Accuracy": [data.Accuracy]
        })

        predicted_cluster = model.predict(input_df)[0]

        return {"WPM": data.WPM, "Accuracy": data.Accuracy, "Predicted_Cluster": int(predicted_cluster)}
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
