import 'package:snake/snake.dart' as snake;

import 'dart:io';
import 'dart:async';

void printMenu() {

}

void main(List<String> arguments) {

  if (!stdin.hasTerminal) {
    stderr.writeln('No terminal attached to stdin.');
    exit(1);
  }

  stdin.echoMode = false;
  stdin.lineMode = false; // NOTE: без нажатия Enter

  ProcessSignal.sigint.watch().listen((_) {
    print("Received SIGINT (Ctrl+C)");
    print("Restoring terminal settings...");

    stdin.echoMode = true;
    stdin.lineMode = true;

    exit(0);
  });

  stdin.listen((List<int> data) {
    for (final byte in data) {
      final ch = String.fromCharCode(byte);
      if (ch == 'w') print('UP');
      if (ch == 's') print('DOWN');
      if (ch == 'a') print('LEFT');
      if (ch == 'd') print('RIGHT');

      if (byte == 113) exit(0); // 'q'
    }

    print("End for...");
  });

  List<List<String>> gameBoard = [
    ['#', '#', '#', '#', '#'],
    ['#', ' ', ' ', ' ', '#'],
    ['#', ' ', '@', ' ', '#'],
    ['#', ' ', ' ', ' ', '#'],
    ['#', '#', '#', '#', '#'],
  ];

  Timer.periodic(Duration(milliseconds: 500), (timer) {
    print('\x1B[2J\x1B[H');

    for (var row in gameBoard) {
      print(row.join(''));
    }

    // Можно остановить таймер по условию
    // if (someCondition) timer.cancel();
  });


}
