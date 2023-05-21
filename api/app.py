from fastapi import FastAPI
from database.models import create_all

app = FastAPI()

create_all()