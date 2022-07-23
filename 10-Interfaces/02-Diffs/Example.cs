public interface IPerson
{
    string Name { get; set; }
    int Age { get; set; }
    void Print()
    {
        Console.WriteLine($"{Name} {Age}");
    }
}

public class Person : IPerson
{
    public string Name { get; set; }
    public int Age { get; set; }
    public void Print()
    {
        Console.WriteLine("{0} - {1}", Name, Age);
    }
}