module p_DM_tb  ; 
 
  reg    we   ; 
  reg  [31:0]  in_d   ; 
  reg    clk   ; 
  wire  [31:0]  out_d   ; 
  reg  [11:2]  addr   ; 
  initial
  begin
    we = 0;
    in_d = 0;
    clk = 0;
    addr = 0;
  end
  always #1 clk = ~clk;
  always #4 we = ~we;
  always #4 in_d = in_d + 8;
  always #4 addr = addr + 4;
  DM  
   DUT  ( 
       .we (we ) ,
      .in_d (in_d ) ,
      .clk (clk ) ,
      .out_d (out_d ) ,
      .addr (addr ) ); 

endmodule

