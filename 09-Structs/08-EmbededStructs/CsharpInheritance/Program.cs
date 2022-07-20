namespace CsharpInheritance
{
    public class Program
    {
        public static void Main(string[] args)
        {
            if (args.Length > 0 && args?[0] == "good")
                Good.Example.Call();
            else
                Bad.Example.Call();

        }
    }
}