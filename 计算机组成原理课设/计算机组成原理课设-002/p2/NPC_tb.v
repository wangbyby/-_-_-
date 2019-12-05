module NPC_tb  ; 
 
  reg    beq   ; 
  reg    jr   ; 
  wire  [31:0]  npc   ; 
  reg    flag_zero   ; 
  reg  [31:0]  gpr_rs   ; 
  wire  [31:0]  jal_out   ; 
  reg    j   ; 
  reg    jal   ; 
  reg  [15:0]  target   ; 
  reg  [25:0]  instr_index   ; 
  reg  [31:0]  pc   ; 
  initial
  begin
    beq = 0;
    jr = 0;
    j =0;
    jal = 0;
    flag_zero = 1;
    gpr_rs = 0;
    target = 0;
    instr_index = 0;
    pc = 0;
  end
   always #1 gpr_rs = gpr_rs + 1;
  always #1 pc = pc +1;
  always #2 flag_zero = flag_zero +  1;
  always #1 target = target +1;
  
  always #1 {beq,j,jal,jr} = {beq,j,jal,jr} +1;
  NPC  
   DUT  ( 
       .beq (beq ) ,
      .jr (jr ) ,
      .npc (npc ) ,
      .flag_zero (flag_zero ) ,
      .gpr_rs (gpr_rs ) ,
      .jal_out (jal_out ) ,
      .j (j ) ,
      .jal (jal ) ,
      .target (target ) ,
      .instr_index (instr_index ) ,
      .pc (pc ) ); 

endmodule

