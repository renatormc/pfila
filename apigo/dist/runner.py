import subprocess
from pathlib import Path
import sys

path = Path(sys.argv[1])
with path.open('w') as f:
    try:
        subprocess.check_call(sys.argv[2:],  stdout=f, stderr=subprocess.STDOUT)
        f.write(f"\n#pfilaok#")
    except:
        f.write(f"\n#pfilaerror#")
