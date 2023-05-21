from sqlalchemy.orm import DeclarativeBase, Session
from sqlalchemy import create_engine
import config


class Base(DeclarativeBase):
    pass


engine = create_engine(config.SQLALCHEMY_DATABASE_URI)


def get_db():
    db = Session(engine)
    try:
        yield db
    finally:
        db.close()



