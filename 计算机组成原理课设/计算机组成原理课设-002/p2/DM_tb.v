module DM_tb  ; 
  reg read;
  reg    we   ; 
  reg  [31:0]  in_d   ; 
  reg    clk   ; 
  wire  [31:0]  out_d   ; 
  reg  [11:2]  addr   ; 
  initial
  begin
    we = 0;
    clk = 0;
    addr = 10;
    in_d = 8;
    read = 0;
    #1000
    $stop();
  end
  always #1 clk = ~clk;
  always #2 we = ~we;
  always #4 read = ~read;
  always #4 in_d = in_d ;
  always #4 addr = addr ;
  
  DM  
   DUT  ( 
       .we (we ) ,
       .read(read ),
      .in_d (in_d ) ,
      .clk (clk ) ,
      .out_d (out_d ) ,
      .addr (addr ) ); 

endmodule

