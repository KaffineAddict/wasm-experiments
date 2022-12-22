namespace HelloDotnet;
public class Library
{
    [System.Runtime.InteropServices.UnmanagedCallersOnly(EntryPoint = "hello")]
    public static int Hello()
    {

        return 42;
    }

    [System.Runtime.InteropServices.UnmanagedCallersOnly(EntryPoint = "add")]
    public static int Add(int x, int y)
    {
        return x+y;
    }
}

//dotnet publish /p:NativeLib=Static /p:SelfContained=true -r browser-wasm -c Debug /p:TargetArchitecture=wasm /p:PlatformTarget=AnyCPU /p:MSBuildEnableWorkloadResolver=false /p:EmccExtraArgs="-s EXPORTED_FUNCTIONS=_add%2C_hello -s EXPORTED_RUNTIME_METHODS=cwrap" --self-contained