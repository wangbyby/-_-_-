module DM(
  input clk,
  input we, //ram write enable
  input read, // ram read enable
  input  [11:2] addr,
  input  [31:0] in_d,
  output [31:0]  out_d );
//4k
  reg [31:0]ram[1023:0];
integer i;
initial begin
  for(i=0;i<1023;i=i+1)
    ram[i] = 0;
end
  assign out_d = (read)? ram[addr] : 32'b0;
  always@(posedge clk) begin
    if(we == 1) ram[addr] <= in_d;
  end
endmodule