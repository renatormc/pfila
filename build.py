from pathlib import Path
import os 

workdir = Path(".").absolute()
os.chdir(workdir / "api")
print("Compilando API...")
if os.name == "nt":
    os.system(".\\build.bat")
else:
    os.system("./build")
print("Compilando interface...")
os.chdir(workdir / "interface")
os.environ['ENV'] = 'prod'
os.system("npx vite build --emptyOutDir")