	 .arch armv8-a
	 .text
//fact
	 .type fact, %function
	 .global fact
	 .p2align 2
fact:
// Prologue
	 sub sp, sp, #112
	 stp x29, x30, [sp]
	 mov x29, sp
.L0:
//store i64 %x, i64* %P_x

	str x0, [sp, #96]
//br label %L1

	b .L1
.L1:
//%r0 = load i64, i64* %P_x

	ldr x8, [sp, #96]
	str x8, [sp, #24]
//%r1 = icmp sle i64 %r0, 1

	ldr x8, [sp, #24]
	mov x9, #1
	cmp x8, x9
	b.le .L2
	b .L3

.L2:
//store i64 1, i64* %fact_retval

	mov x8, #1
	str x8, [sp, #16]
//%r2 = load i64, i64* %fact_retval

	ldr x8, [sp, #16]
	str x8, [sp, #40]
//ret i64 %r2

	ldr x0, [sp, #%r2]
.L3:
//%r3 = load i64, i64* %P_x

	ldr x8, [sp, #96]
	str x8, [sp, #48]
//%r4 = load i64, i64* %P_x

	ldr x8, [sp, #96]
	str x8, [sp, #56]
//%r5 = sub i64 %r4, 1

	ldr x8, [sp, #56]
	mov x9, #1
	sub x10, x8, x9
	str x10, [sp, #64]
//%r6 = call i64 @fact(i64 %r5)

	ldr x0, [sp #64]
	bl fact
	str x0, [sp, #72]
//%r7 = mul i64 %r3, %r6

	ldr x8, [sp, #48]
	ldr x9, [sp, #72]
	mul x10, x8, x9
	str x10, [sp, #80]
//store i64 %r7, i64* %fact_retval

	ldr x8, [sp, #80]
	str x8, [sp, #16]
//%r8 = load i64, i64* %fact_retval

	ldr x8, [sp, #16]
	str x8, [sp, #88]
//ret i64 %r8

	ldr x0, [sp, #%r8]
// Epilogue
	 ldp x29, x30, [sp]
	 add sp, sp, #112
	 ret
	 .size fact, (.-fact)

//main
	 .type main, %function
	 .global main
	 .p2align 2
main:
// Prologue
	 sub sp, sp, #96
	 stp x29, x30, [sp]
	 mov x29, sp
.L4:
//store i64 0, i64* %stop

	mov x8, #0
	str x8, [sp, #32]
//store i64 0, i64* %factor

	mov x8, #0
	str x8, [sp, #40]
//call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %factor)

	adrp x8, .READ
	add x8, x8, :lo12:.READ
	mov x0, x8
	add x1, sp,  #40
	bl scanf
//%r9 = load i64, i64* %factor

	ldr x8, [sp, #40]
	str x8, [sp, #56]
//%r10 = call i64 @fact(i64 %r9)

	ldr x0, [sp #56]
	bl fact
	str x0, [sp, #64]
//store i64 %r10, i64* %temp

	ldr x8, [sp, #64]
	str x8, [sp, #24]
//%r11 = load i64, i64* %temp

	ldr x8, [sp, #24]
	str x8, [sp, #72]
//call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.fstr1, i32 0, i32 0), i64 %r11)

	adrp x8, .FMTfstr1
	add x8, x8, :lo12:.FMTfstr1
	mov x0, x8
	ldr x1, [sp, #72]
	bl printf
//store i64 0, i64* %main_retval

	mov x8, #0
	str x8, [sp, #16]
//%r12 = load i64, i64* %main_retval

	ldr x8, [sp, #16]
	str x8, [sp, #80]
//ret i64 %r12

	ldr x0, [sp, #%r12]
// Epilogue
	 ldp x29, x30, [sp]
	 add sp, sp, #96
	 ret
	 .size main, (.-main)

.READ:
	 .asciz "%ld\00"
	 .size .READ, 4
.FMTfstr1:
	 .asciz "%ld\00"
	 .size .FMTfstr1, 4
