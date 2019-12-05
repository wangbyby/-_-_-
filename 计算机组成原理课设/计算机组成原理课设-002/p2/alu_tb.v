module alu_tb  ; 
 
  reg  [31:0]  inA   ; 
  wire   flag_add_overflow   ; 
  wire  flag_beq   ; 
  reg  [3:0]  sel_alu   ; 
  reg  [31:0]  inB   ; 
  reg  [15:0]  imm16   ; 
  reg    clk   ; 
  wire  [31:0]  outC   ; 
  initial
  begin
    clk = 0;
    inA = 0;
    inB = 0;
    sel_alu = 0;
    imm16 = 0;
    #1000
    
    $stop();
  end
  always #2 clk = ~clk;
  always #6 inA = inA + 1;
  always #6 inB = inB + 1;
  always #6 sel_alu = sel_alu + 1;
  always #6 imm16 = imm16 +1;
  alu  
   DUT  ( 
       .inA (inA ) ,
      .flag_add_overflow (flag_add_overflow ) ,
      .flag_beq (flag_beq ) ,
      .sel_alu (sel_alu ) ,
      .inB (inB ) ,
      .imm16 (imm16 ) ,
      .clk (clk ) ,
      .outC (outC ) ); 

endmodule

