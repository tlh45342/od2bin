.globl  _main
.text
_main:
jsr     r5,rsave; 2
.data
L2:.byte 150,145,154,154,157,12,0
.even
.text
mov     $6,(sp)
mov     $L2,-(sp)
mov     $1,-(sp)
jsr     pc,*$_write
cmp     (sp)+,(sp)+
L1:jmp  rretrn
