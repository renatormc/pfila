from database.db import Base
from sqlalchemy.orm import mapped_column, Mapped
import sqlalchemy as sa
from datetime import datetime
import json

class Process(Base):
    __tablename__ = 'process'
    id: Mapped[int] = mapped_column(sa.Integer, primary_key=True)
    type: Mapped[str] = mapped_column(sa.String(150))
    name: Mapped[str] = mapped_column(sa.String(300))
    user: Mapped[str] = mapped_column(sa.String(300))
    pid: Mapped[int | None] = mapped_column(sa.Integer)
    created_at: Mapped[datetime | None] = mapped_column(sa.DateTime)
    started_at: Mapped[datetime | None] = mapped_column(sa.DateTime)
    started_waiting_at: Mapped[datetime | None] = mapped_column(sa.DateTime)
    finished_at: Mapped[datetime | None] = mapped_column(sa.DateTime)
    status: Mapped[str] = mapped_column(sa.String(150))
    random_id: Mapped[str] = mapped_column(sa.String(150))
    params: Mapped[str] = mapped_column(sa.Text)
    dependencies: Mapped[str] = mapped_column(sa.String(500))
    is_docker: Mapped[bool] = mapped_column(sa.Boolean, nullable=False, default=False)

    def set_params(self, value) -> None:
        self.params = json.dumps(value)

    def get_params(self) -> dict:
        return json.loads(self.params)
    
    def get_dependencies(self) -> list[int]:
        if self.dependencies == "":
            return []
        text = self.dependencies[1:-1]
        try:
            return [int(p) for p in text.split(",")]
        except ValueError:
            return []
        
    def set_dependencies(self, deps: list[int]) -> None:
        text = ",".join([str(d) for d in deps])
        self.dependencies = f",{text},"
    