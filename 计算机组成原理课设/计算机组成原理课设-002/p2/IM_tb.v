module IM_tb  ; 
 
  wire  [31:0]  outIns   ; 
  reg  [31:0]  address   ; 
  initial begin
    address = 0;
    #1000
    $stop();
  end
  always #1 address = address + 4;
  IM  
   DUT  ( 
       .outIns (outIns ) ,
      .address (address ) ); 

endmodule

