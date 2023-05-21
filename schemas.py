from pydantic import BaseModel
from datetime import datetime
from typing import Optional

class ProcessCreateSchema(BaseModel):
    type: str
    name: str
    user: str
    pid: Optional[int]
    created_at: Optional[datetime]
    started_at: Optional[datetime]
    finished_at: Optional[datetime]
    status: str
    random_id: str
    params: str
    dependencies: str


class ProcessSchema(ProcessCreateSchema):
    id: int

    class Config:
        orm_mode = True
