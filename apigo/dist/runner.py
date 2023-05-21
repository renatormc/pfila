import subprocess
from pathlib import Path

args = ["D:\\tests\\pfila\\iped\\iped-4.1.2\\jre\\bin\\java.exe", "-jar",
        "D:\\tests\\pfila\\iped\\iped-4.1.2\\iped.jar", "-profile", "fastmode",
        "-d", "D:\\tests\\pfila\\pen.E01", "-o", "D:\\tests\\pfila\\result", "--nogui"]
path = Path(r'D:\tests\pfila\console\output.txt')
with path.open('w') as f:
    subprocess.run(args,  stdout=f, stderr=subprocess.STDOUT)
