from database.db import engine
from .models import Process, Base

def create_all():
    Base.metadata.create_all(engine)