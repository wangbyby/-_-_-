module pcCal( clk,imm16,target26,pc,flag_mod_select,npc,pc_add_4);
input clk;
input [15:0] imm16;
input [29:0] pc;
input [25:0] target26;
input [2:0] flag_mod_select;
output [29:0] npc,pc_add_4;

always @(posedge clk or flag_mod_select) begin
    case (flag_mod_select)
        : // npc = pc + 4

        :  // beq 
            npc = pc + 1'b1 + imm16[15:2];
        end
        : //jal 
        : //jr
        : //j
        default: // npc = pc + 4
    endcase
end
assign pc_add_4 = pc+1'b1;
endmodule // pcCal 