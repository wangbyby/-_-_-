module PC_tb  ; 
 
  reg    pcReset   ; 
  reg    clk   ; 
  wire  [31:0]  pcOut   ; 
  reg  [31:0]  pc   ;
  initial
  begin
    clk = 0;
    pcReset = 0;
    pc = 0;
    #1000
    $stop();
end 

always #1 clk = ~clk;
always #1 pc = pc + 4;
always #5 pcReset = ~pcReset;

  PC  
   DUT  ( 
       .pcReset (pcReset ) ,
      .clk (clk ) ,
      .pcOut (pcOut ) ,
      .pc (pc ) ); 

endmodule

