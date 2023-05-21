from pathlib import Path
import os
import tomllib

APPDIR = Path(os.path.dirname(os.path.realpath(__file__)))

path = APPDIR / "pfila.toml"
with path.open("rb") as f:
    data = tomllib.load(f)

CONSOLE_FOLDER: str = data['console_folder']
PORT: str = data['port']
CHECK_AUTH: bool = data['check_auth']
SECRET: str = data['secret']
IPED_FOLDER = Path(data['iped_folder'])
IPED_PROFILE_FOLDER = Path(data['iped_profile_folder'])
p = APPDIR / "pfile.db"
SQLALCHEMY_DATABASE_URI= f"sqlite:///{p}"