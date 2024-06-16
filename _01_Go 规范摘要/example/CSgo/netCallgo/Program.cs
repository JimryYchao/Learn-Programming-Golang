using System.Runtime.InteropServices;

unsafe
{
    Console.WriteLine("Hello, World!");
    Sample.Hello();
    GoSlice<long> s = new long[] { 1,2,3,4,5,6 };
    //Sample.Println("%d", s);

    Sample.IterIntSlice(s);
}

class Sample
{
    [DllImport("CSgo.Interop.dll", EntryPoint = "Hello", CallingConvention = CallingConvention.Cdecl)]
    public static extern void Hello();

    [DllImport("CSgo.Interop.dll", EntryPoint = "Println", CallingConvention = CallingConvention.Cdecl)]
    public static unsafe extern void Println(GoString format, GoSlice args);

    [DllImport("CSgo.Interop.dll", EntryPoint = "IterIntSlice", CallingConvention = CallingConvention.Cdecl)]
    public extern static void IterIntSlice(GoSlice<long> slice);
}


internal struct GoSlice<T> : IDisposable
{
    public IntPtr data;
    public long len, cap;
    public GoSlice(IntPtr data, long len, long cap)
    {
        this.data = data;
        this.len = len;
        this.cap = cap;
    }

    public void Dispose()
    {
        Marshal.FreeHGlobal(data);
    }

    public unsafe static implicit operator GoSlice<T>(T[] data)
    {
        var data_ptr = Marshal.AllocHGlobal(Buffer.ByteLength(data));
        fixed (void* ptr = &data[0])
        {
            Buffer.MemoryCopy(ptr, (void*)data_ptr, sizeof(T) / sizeof(byte) * data.Length, sizeof(T) / sizeof(byte) * data.Length);
        }
        return new(data_ptr, data.Length, data.Length);
    }

    public static implicit operator GoSlice(GoSlice<T> slice)
    {
        return new GoSlice(slice.data, slice.len, slice.cap);
    }
}

internal struct GoSlice:IDisposable
{
    public IntPtr data;
    public long len, cap;
    public GoSlice(IntPtr data, long len, long cap)
    {
        this.data = data;
        this.len = len;
        this.cap = cap;
    }

    public void Dispose()
    {
        Marshal.FreeHGlobal(data);
    }
}
internal struct GoString
{
    public string msg;
    public long len;
    public GoString(string msg, long len)
    {
        this.msg = msg;
        this.len = len;
    }

    public static implicit operator GoString(string msg)
    {
        return new GoString(msg, msg.Length);
    }

}