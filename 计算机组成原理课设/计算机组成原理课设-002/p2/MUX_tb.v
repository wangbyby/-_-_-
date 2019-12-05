module MUX_tb_sec  ; 
 
  reg    choose   ; 
  reg  [31:0]  a   ; 
  wire  [31:0]  z   ; 
  reg  [31:0]  b   ; 
  initial
  begin
    a = 0;
    b = 0;
    choose = 0;
end
  always #1 a = a +1;
  always #2 b = b + 1;
  always #3 choose = choose + 1;
  MUX  
   DUT  ( 
       .choose (choose ) ,
      .a (a ) ,
      .z (z ) ,
      .b (b ) ); 

endmodule

