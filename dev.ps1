& $PSScriptRoot\env.ps1
$WORKDIR = Get-Location
try {
    switch ($args[0]) {
        "api" {
            Set-Location "$WORKDIR\api"
            go build -o "$WORKDIR\dist\pfila.exe"
            Set-Location "$WORKDIR\dist"
            Invoke-Expression ".\pfila.exe serve"
        }
        "interface" {
            Set-Location "$WORKDIR\interface"
            yarn dev
        }
        "build" {
            Set-Location "$WORKDIR\api"
            go build -o "$WORKDIR\dist\pfila.exe"
            Write-Host "Compiling interface..."
            Set-Location "$WORKDIR\interface"
            $env:ENV = "prod"
            npx vite build --emptyOutDir
        }
        default {
            Write-Host "Command unknown"
        }
    }
}
finally {
    Set-Location $WORKDIR
}

