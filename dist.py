from pathlib import Path
import os
import shutil

folder = Path(".").absolute()

os.chdir(folder / "api")
os.system(".\\build.bat")
os.chdir(folder / "interface")
os.system("npx vite build")

dist_folder = folder / "dist"
try:
    shutil.rmtree(dist_folder)
except FileNotFoundError:
    pass
dist_folder.mkdir()

shutil.copy(folder / "api/dist/pfila.exe", dist_folder / "pfila.exe")
shutil.copy(folder / "api/dist/pfila.toml", dist_folder / "pfila.toml")
shutil.copytree(folder / "api/dist/app", dist_folder / "app")
