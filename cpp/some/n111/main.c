/* Инициализированная глобальная переменная */
int z_global = 11;
/* Вторая глобальная переменная с именем y_global_init, но они обе static */
static int y_global_init = 2;
/* Объявление другой глобальной переменной */
extern int x_global_init;

int fn_a(int x, int y)
{
    return(x+y);
}

int x; // .bss
int *addr = &x;

int main(int argc, char *argv)
{
    printf("%d\n", addr);
    x = 10;
    printf("%d\n", &x);

    const char *message = "Hello, world";
    return fn_a(11,12);
}
