from typing import List
from .app import app
from fastapi import Depends, HTTPException
from database.db import get_db
from sqlalchemy.orm import Session
from schemas import ProcessCreateSchema, ProcessSchema
from database.models import Process
import config

@app.get("/api/proc", response_model=List[ProcessSchema])
def list_procs(db: Session = Depends(get_db)):
    procs = db.query(Process).all()
    return procs


@app.get("/api/proc/{process_id}", response_model=ProcessSchema)
def get_process(process_id: int, db: Session = Depends(get_db)):
    proc = db.query(Process).filter(Process.id == process_id).first()
    if not proc:
        raise HTTPException(status_code=404, detail="Process not found")
    return proc


@app.post("/api/proc",  response_model=ProcessSchema)
async def create_proc(process: ProcessCreateSchema, db: Session = Depends(get_db)):
    proc = Process(**process.dict())
    db.add(proc)
    db.commit()
    return proc


@app.put("/api/proc/{process_id}", response_model=ProcessSchema)
def update_process(process_id: int, process: ProcessCreateSchema, db: Session = Depends(get_db)):
    proc = db.query(Process).filter(Process.id == process_id).first()
    if not proc:
        raise HTTPException(status_code=404, detail="Process not found")

    for key, value in process.dict().items():
        setattr(proc, key, value)

    db.commit()
    return proc


@app.delete("/api/proc/{process_id}")
def delete_process(process_id: int, db: Session = Depends(get_db)):
    proc = db.query(Process).filter(Process.id == process_id).first()
    if not proc:
        raise HTTPException(status_code=404, detail="Process not found")
    db.delete(proc)
    db.commit()
    return {"message": "Process deleted successfully"}


@app.get("/api/iped-profiles")
def iped_profiles():
    return [entry.name for entry in config.IPED_PROFILE_FOLDER.iterdir() if entry.is_dir()]