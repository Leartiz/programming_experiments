import 'package:hello/hello.dart' as hello;

void main(List<String> arguments) {
  print('Hello world: ${hello.calculate()}!');

  {
    print("Hello!");
    print("Welcome to Dart!");
  }

  // var
  {
    var name = "Tom";
    print(name);

    //name = 45;

    print(name);
  }

  // dynamic
  {
    dynamic name = "Tom";
    print(name);
    print("name.runtimeType: ${name.runtimeType}");

    name = 45;

    print(name);
    print("name.runtimeType: ${name.runtimeType}");
  }

  // final
  {
    final DateTime today = DateTime.now();
    print(today);
    //today = DateTime.now();
  }

  // String
  {
    var multiline = '''
Многострочная
строка
    ''';
    print(multiline);
    String text = """
Высокой страсти не имея
Для кода жизни не щадить,
Не мог он джаву от сишарпа,
Как мы ни бились, отличить.
    """;
    print(text);
  }

  // interpolation
  {
    String name = "Tom";
    int age = 35;

    String info = "Name: $name  Age: $age";
    print("info: $info");
  }

  // null
  {
    String? name = "Tom";
    print(name);
    name = null;
    print(name);
  }

  // operator ??
  {
    int? num1 = 23;
    int num2 = num1 ?? 0;
    print(num2);

    num1 = null;
    num2 = num1 ?? 0;
    print(num2);

    int? a;
    a ??= 10;
    print("a: $a");
  }

  // if-else
  {
    const int num1 = 8;
    const int num2 = 5;
    if (num1 > num2) {
      print("Первое число больше второго");
    } else if (num1 < num2) {
      print("Первое число меньше второго");
    } else {
      print("Числа равны");
    }
  }

  // for
}
