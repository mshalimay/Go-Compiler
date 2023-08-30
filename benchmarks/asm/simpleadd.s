	 .arch armv8-a
	 .text
//main
	 .type main, %function
	 .global main
	 .p2align 2
main:
// Prologue
	 sub sp, sp, #112
	 stp x29, x30, [sp]
	 mov x29, sp
.L0:
//store i64 129, i64* %a

	mov x8, #129
	str x8, [sp, #24]
//call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %b)

	adrp x8, .READ
	add x8, x8, :lo12:.READ
	mov x0, x8
	add x1, sp,  #32
	bl scanf
//%r0 = load i64, i64* %a

	ldr x8, [sp, #24]
	str x8, [sp, #48]
//%r1 = load i64, i64* %b

	ldr x8, [sp, #32]
	str x8, [sp, #56]
//%r2 = add i64 %r0, %r1

	ldr x8, [sp, #48]
	ldr x9, [sp, #56]
	add x10, x8, x9
	str x10, [sp, #64]
//store i64 %r2, i64* %c

	ldr x8, [sp, #64]
	str x8, [sp, #40]
//%r3 = load i64, i64* %a

	ldr x8, [sp, #24]
	str x8, [sp, #72]
//call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([5 x i8], [5 x i8]* @.fstr1, i32 0, i32 0), i64 %r3)

	adrp x8, .FMTfstr1
	add x8, x8, :lo12:.FMTfstr1
	mov x0, x8
	ldr x1, [sp, #72]
	bl printf
//%r4 = load i64, i64* %b

	ldr x8, [sp, #32]
	str x8, [sp, #80]
//call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([5 x i8], [5 x i8]* @.fstr2, i32 0, i32 0), i64 %r4)

	adrp x8, .FMTfstr2
	add x8, x8, :lo12:.FMTfstr2
	mov x0, x8
	ldr x1, [sp, #80]
	bl printf
//%r5 = load i64, i64* %c

	ldr x8, [sp, #40]
	str x8, [sp, #88]
//call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.fstr3, i32 0, i32 0), i64 %r5)

	adrp x8, .FMTfstr3
	add x8, x8, :lo12:.FMTfstr3
	mov x0, x8
	ldr x1, [sp, #88]
	bl printf
//store i64 0, i64* %main_retval

	mov x8, #0
	str x8, [sp, #16]
//%r6 = load i64, i64* %main_retval

	ldr x8, [sp, #16]
	str x8, [sp, #96]
//ret i64 %r6

	ldr x0, [sp, #%r6]
// Epilogue
	 ldp x29, x30, [sp]
	 add sp, sp, #112
	 ret
	 .size main, (.-main)

.READ:
	 .asciz "%ld\00"
	 .size .READ, 4
.FMTfstr1:
	 .asciz "%ld\00"
	 .size .FMTfstr1, 5
.FMTfstr2:
	 .asciz "%ld\00"
	 .size .FMTfstr2, 5
.FMTfstr3:
	 .asciz "%ld\00"
	 .size .FMTfstr3, 4
