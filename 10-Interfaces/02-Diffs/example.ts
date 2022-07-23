interface IPerson {
    name: string;
    age: number;
    print();

}
 
class Person implements IPerson {
    name: string;
    age: number; 
    print() {
        console.log(this.name);
    }
}