using System;
using Newtonsoft;
using Newtonsoft.Json;

namespace Bad
{
    public class Example
    {
        public static void Call()
        {

            var mobile = new Mobile { Name = "Samsung M510" };
            mobile.Brand = "Samsung";
            mobile.Color = "Blue";
            mobile.Ram = 4;
            mobile.OperatingSystem = "Android";
            mobile.OperatingSystemVersion = "11";

            mobile.CameraCount = 2;
            mobile.NetworkType = "5G";
            mobile.SimcardType = "Nano";
            mobile.SimcardCapacity = 3;

            var loptop = new Loptop { Name = "MSI GX4450" };

            loptop.Brand = "MSI";
            loptop.Color = "Black";
            loptop.Ram = 16;
            loptop.OperatingSystem = "Ubuntu";
            loptop.OperatingSystemVersion = "22.04";

            loptop.HasCdRom = false;
            loptop.UsbPortCount = 3;
            loptop.KeyboardType = "Light";

            Console.WriteLine(JsonConvert.SerializeObject(mobile));
            Console.WriteLine(JsonConvert.SerializeObject(loptop));

        }
    }

    ////////////////////////////////////////////

    public class Mobile
    {
        public string Name { get; set; }
        public string Color { get; set; }
        public float Length { get; set; }
        public float Width { get; set; }
        public float Weight { get; set; }
        public decimal Price { get; set; }
        public string Brand { get; set; }
        public string MadeIn { get; set; }
        public int Ram { get; set; }
        public string Cpu { get; set; }
        public float ScreenSize { get; set; }
        public string OperatingSystem { get; set; }
        public string OperatingSystemVersion { get; set; }
        public int SimcardCapacity { get; set; }
        public string SimcardType { get; set; }
        public string NetworkType { get; set; }
        public int CameraCount { get; set; }
    }

    public class Loptop
    {
        public string Name { get; set; }
        public string Color { get; set; }
        public float Length { get; set; }
        public float Width { get; set; }
        public float Weight { get; set; }
        public decimal Price { get; set; }
        public string Brand { get; set; }
        public string MadeIn { get; set; }
        public int Ram { get; set; }
        public string Cpu { get; set; }
        public float ScreenSize { get; set; }
        public string OperatingSystem { get; set; }
        public string OperatingSystemVersion { get; set; }
        public int UsbPortCount { get; set; }
        public string KeyboardType { get; set; }
        public bool HasCdRom { get; set; }
    }
}