using System;

namespace HelloWorld
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Hello World!");
        }

        [System.Runtime.InteropServices.UnmanagedCallersOnly(EntryPoint = "hello")]
        public static int Hello()
        {
            return 42;
        }

        [System.Runtime.InteropServices.UnmanagedCallersOnly(EntryPoint = "add")]
        public static int Add(int x, int y)
        {
            return x + y;
        }
    }
}