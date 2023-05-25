from pathlib import Path
import os 

workdir = Path(".").absolute()
os.chdir(workdir / "api")
print("Compilando API...")
os.system("build.bat")
print("Compilando interface...")
os.chdir(workdir / "interface")
os.system("npx vite build --emptyOutDir")