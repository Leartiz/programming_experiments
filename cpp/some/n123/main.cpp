#include <stdlib.h>
#include <setjmp.h>
#include <unistd.h>
#include <iostream>

// -----------------------------------------------------------------------

typedef struct {
    jmp_buf context; // Контекст для сохранения состояния
    void (*function)(void); // Указатель на функцию "файбера"
} Fiber;

static Fiber
    fiber1, fiber2;

void fiber1_function(void);
void fiber2_function(void);

void fiber_switch(Fiber *current, Fiber *next) {
    if (setjmp(current->context) == 0) {
        longjmp(next->context, 1);
    }
}

// -----------------------------------------------------------------------

void fiber1_function(void) {
    for (int i = 0; i < 3; i++) {
        std::cout << "Fiber 1: " << i << '\n';
        sleep(1);

        fiber_switch(&fiber1, &fiber2);
    }
}

void fiber2_function(void) {
    for (int i = 0; i < 3; i++) {
        std::cout << "Fiber 2: " << i << '\n';
        sleep(1);

        fiber_switch(&fiber2, &fiber1);
    }
}

int main() {

    // Устанавливаем указатели на функции
    fiber1.function = fiber1_function;
    fiber2.function = fiber2_function;

    // Устанавливаем контекст для первого файбера
    if (setjmp(fiber1.context) == 0) {
        fiber1.function();
    }

    return 0;
}
