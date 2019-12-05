module Controll_tb  ; 
 
  wire   beq   ; 
  wire  lui   ; 
  wire  [3:0]  sel_ALU   ; 
  wire   rd   ; 
  reg  [5:0]  op   ; 
  wire    Extop   ; 
  wire    RAM_to_GPR   ; 
  wire   RAM_write   ; 
  wire    j   ; 
  wire    GPR_write   ; 
  wire    jr   ; 
  wire    imm_to_ALU   ; 
  reg  [5:0]  order_func   ; 
  wire  jal   ; 
  initial
  begin
    op = 0;
    order_func = 0;
    #2000
    $stop();
  end
  always #1 op = op+1;
  always #1 order_func = order_func + 1;
  Controll  
   DUT  ( 
       .beq (beq ) ,
      .lui (lui ) ,
      .sel_ALU (sel_ALU ) ,
      .rd (rd ) ,
      .op (op ) ,
      .Extop (Extop ) ,
      .RAM_to_GPR (RAM_to_GPR ) ,
      .RAM_write (RAM_write ) ,
      .j (j ) ,
      .GPR_write (GPR_write ) ,
      .jr (jr ) ,
      .imm_to_ALU (imm_to_ALU ) ,
      .order_func (order_func ) ,
      .jal (jal ) ); 

endmodule

