`define WIDTH 32
`define DEPTH 1024
`define ADDR_SIZE 10 

module DM(clk,rst_n,rd_en,wr_en,rdaddr,data_in,data_out);
	input clk,rst_n;
	input rd_en,wr_en;
	input [`ADDR_SIZE-1:0] rdaddr;
	input [`WIDTH-1:0] data_in;
	
	reg [`WIDTH-1:0]mem[`DEPTH-1:0];
	output [`WIDTH-1:0] data_out;
	reg [`WIDTH-1:0] data_out;
	

always @(posedge clk or negedge rst_n)begin
	if(rst_n)
		data_out<=0;
	else if(rd_en)
		data_out <= mem[rdaddr];
	else if(wr_en)
		mem[rdaddr] <= data_in;	
end
endmodule
