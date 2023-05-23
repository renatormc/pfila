import subprocess
from pathlib import Path
import sys
from database.db import engine
from sqlalchemy.orm import Session
from sqlalchemy import select
from database.models import Process


def get_next_proc() -> Process | None:
    with Session(engine) as session:
        stm = select(Process).where(Process.status == "AGUARDANDO").order_by(Process.started_waiting_at.asc())
        return session.execute(stm).scalars().first()


path = Path(sys.argv[1])
with path.open('w') as f:
    try:
        subprocess.check_call(sys.argv[2:],  stdout=f, stderr=subprocess.STDOUT)
        f.write(f"\n#pfilaok#")
    except:
        f.write(f"\n#pfilaerror#")
