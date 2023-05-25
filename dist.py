from pathlib import Path
import os
import shutil

folder = Path(".").absolute()

os.chdir(folder / "api")
os.system(".\\build.bat")
os.chdir(folder / "interface")
os.system("npx vite build")

dist_folder = folder / "dist"
for entry in dist_folder.iterdir():
    if entry.name == "tools":
        continue
    try:
        if entry.is_dir():
            shutil.rmtree(entry)
        else:
            entry.unlink()
    except FileNotFoundError:
        pass

shutil.copy(folder / "api/dist/pfila.exe", dist_folder / "pfila.exe")
shutil.copy(folder / "api/dist/update.exe", dist_folder / "update.exe")
shutil.copy(folder / "api/dist/pfila.toml", dist_folder / "pfila.toml")
shutil.copytree(folder / "api/dist/app", dist_folder / "app")
