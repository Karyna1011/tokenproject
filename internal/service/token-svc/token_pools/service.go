package token

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"gitlab.com/tokend/subgroup/tokenproject/internal/data"

	"gitlab.com/tokend/subgroup/tokenproject/internal/config"
	"strconv"
	"strings"
	//"time"
)

const myABI="[{\"name\":\"NewAddressIdentifier\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"id\",\"indexed\":true},{\"type\":\"address\",\"name\":\"addr\",\"indexed\":false},{\"type\":\"string\",\"name\":\"description\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddressModified\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"id\",\"indexed\":true},{\"type\":\"address\",\"name\":\"new_address\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"version\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewAdmin\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"deadline\",\"indexed\":true},{\"type\":\"address\",\"name\":\"admin\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewAdmin\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_admin\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"name\":\"get_registry\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1061},{\"name\":\"max_id\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1258},{\"name\":\"get_address\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_id\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1308},{\"name\":\"add_new_id\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_address\"},{\"type\":\"string\",\"name\":\"_description\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":291275},{\"name\":\"set_address\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_id\"},{\"type\":\"address\",\"name\":\"_address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":182430},{\"name\":\"unset_address\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_id\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":101348},{\"name\":\"commit_transfer_ownership\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_new_admin\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":74048},{\"name\":\"apply_transfer_ownership\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":60125},{\"name\":\"revert_transfer_ownership\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":21400},{\"name\":\"admin\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1331},{\"name\":\"transfer_ownership_deadline\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1361},{\"name\":\"future_admin\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1391},{\"name\":\"get_id_info\",\"outputs\":[{\"type\":\"address\",\"name\":\"addr\"},{\"type\":\"bool\",\"name\":\"is_active\"},{\"type\":\"uint256\",\"name\":\"version\"},{\"type\":\"uint256\",\"name\":\"last_modified\"},{\"type\":\"string\",\"name\":\"description\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":12168}]"
const ABI="[{\"name\":\"PoolAdded\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"indexed\":true},{\"name\":\"rate_method_id\",\"type\":\"bytes\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"PoolRemoved\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_address_provider\",\"type\":\"address\"},{\"name\":\"_gauge_controller\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"find_pool_for_coins\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"find_pool_for_coins\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_n_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}],\"gas\":1521},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[8]\"}],\"gas\":12102},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_underlying_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[8]\"}],\"gas\":12194},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_decimals\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}],\"gas\":7874},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_underlying_decimals\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}],\"gas\":7966},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_rates\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}],\"gas\":36992},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_gauges\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[10]\"},{\"name\":\"\",\"type\":\"int128[10]\"}],\"gas\":20157},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}],\"gas\":16583},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_underlying_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}],\"gas\":162842},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_virtual_price_from_lp_token\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":1927},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_A\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":1045},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_parameters\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"A\",\"type\":\"uint256\"},{\"name\":\"future_A\",\"type\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\"},{\"name\":\"admin_fee\",\"type\":\"uint256\"},{\"name\":\"future_fee\",\"type\":\"uint256\"},{\"name\":\"future_admin_fee\",\"type\":\"uint256\"},{\"name\":\"future_owner\",\"type\":\"address\"},{\"name\":\"initial_A\",\"type\":\"uint256\"},{\"name\":\"initial_A_time\",\"type\":\"uint256\"},{\"name\":\"future_A_time\",\"type\":\"uint256\"}],\"gas\":6305},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_fees\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}],\"gas\":1450},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_admin_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}],\"gas\":36454},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coin_indices\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int128\"},{\"name\":\"\",\"type\":\"int128\"},{\"name\":\"\",\"type\":\"bool\"}],\"gas\":27131},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"estimate_gas_used\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":32004},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"is_meta\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":1900},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_pool_name\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"gas\":8323},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coin_swap_count\",\"inputs\":[{\"name\":\"_coin\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":1951},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coin_swap_complement\",\"inputs\":[{\"name\":\"_coin\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2090},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_pool_asset_type\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":2011},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_pool\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_n_coins\",\"type\":\"uint256\"},{\"name\":\"_lp_token\",\"type\":\"address\"},{\"name\":\"_rate_info\",\"type\":\"bytes32\"},{\"name\":\"_decimals\",\"type\":\"uint256\"},{\"name\":\"_underlying_decimals\",\"type\":\"uint256\"},{\"name\":\"_has_initial_A\",\"type\":\"bool\"},{\"name\":\"_is_v1\",\"type\":\"bool\"},{\"name\":\"_name\",\"type\":\"string\"}],\"outputs\":[],\"gas\":61485845},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_pool_without_underlying\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_n_coins\",\"type\":\"uint256\"},{\"name\":\"_lp_token\",\"type\":\"address\"},{\"name\":\"_rate_info\",\"type\":\"bytes32\"},{\"name\":\"_decimals\",\"type\":\"uint256\"},{\"name\":\"_use_rates\",\"type\":\"uint256\"},{\"name\":\"_has_initial_A\",\"type\":\"bool\"},{\"name\":\"_is_v1\",\"type\":\"bool\"},{\"name\":\"_name\",\"type\":\"string\"}],\"outputs\":[],\"gas\":31306062},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_metapool\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_n_coins\",\"type\":\"uint256\"},{\"name\":\"_lp_token\",\"type\":\"address\"},{\"name\":\"_decimals\",\"type\":\"uint256\"},{\"name\":\"_name\",\"type\":\"string\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_metapool\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_n_coins\",\"type\":\"uint256\"},{\"name\":\"_lp_token\",\"type\":\"address\"},{\"name\":\"_decimals\",\"type\":\"uint256\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_base_pool\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_pool\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[],\"gas\":779731418758},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_pool_gas_estimates\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address[5]\"},{\"name\":\"_amount\",\"type\":\"uint256[2][5]\"}],\"outputs\":[],\"gas\":390460},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_coin_gas_estimates\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address[10]\"},{\"name\":\"_amount\",\"type\":\"uint256[10]\"}],\"outputs\":[],\"gas\":392047},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_gas_estimate_contract\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_estimator\",\"type\":\"address\"}],\"outputs\":[],\"gas\":72629},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_liquidity_gauges\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_liquidity_gauges\",\"type\":\"address[10]\"}],\"outputs\":[],\"gas\":400675},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_pool_asset_type\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_asset_type\",\"type\":\"uint256\"}],\"outputs\":[],\"gas\":72667},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"batch_set_pool_asset_type\",\"inputs\":[{\"name\":\"_pools\",\"type\":\"address[32]\"},{\"name\":\"_asset_types\",\"type\":\"uint256[32]\"}],\"outputs\":[],\"gas\":1173447},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"address_provider\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2048},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"gauge_controller\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2078},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pool_list\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2217},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pool_count\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":2138},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"coin_count\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":2168},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coin\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2307},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_pool_from_lp_token\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2443},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_lp_token\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2473},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"last_updated\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":2288}]"
const ABIV="[{\"name\":\"NewRelease\",\"inputs\":[{\"name\":\"release_id\",\"type\":\"uint256\",\"indexed\":true},{\"name\":\"template\",\"type\":\"address\",\"indexed\":false},{\"name\":\"api_version\",\"type\":\"string\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewVault\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true},{\"name\":\"vault_id\",\"type\":\"uint256\",\"indexed\":true},{\"name\":\"vault\",\"type\":\"address\",\"indexed\":false},{\"name\":\"api_version\",\"type\":\"string\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewExperimentalVault\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true},{\"name\":\"deployer\",\"type\":\"address\",\"indexed\":true},{\"name\":\"vault\",\"type\":\"address\",\"indexed\":false},{\"name\":\"api_version\",\"type\":\"string\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewGovernance\",\"inputs\":[{\"name\":\"governance\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"VaultTagged\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\",\"indexed\":false},{\"name\":\"tag\",\"type\":\"string\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setGovernance\",\"inputs\":[{\"name\":\"governance\",\"type\":\"address\"}],\"outputs\":[],\"gas\":36245},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"acceptGovernance\",\"inputs\":[],\"outputs\":[],\"gas\":37517},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"latestRelease\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"gas\":6831},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"latestVault\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2587},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"newRelease\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\"}],\"outputs\":[],\"gas\":82588},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"newVault\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"guardian\",\"type\":\"address\"},{\"name\":\"rewards\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"newVault\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"guardian\",\"type\":\"address\"},{\"name\":\"rewards\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"releaseDelta\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"newExperimentalVault\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"governance\",\"type\":\"address\"},{\"name\":\"guardian\",\"type\":\"address\"},{\"name\":\"rewards\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"newExperimentalVault\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"governance\",\"type\":\"address\"},{\"name\":\"guardian\",\"type\":\"address\"},{\"name\":\"rewards\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"releaseDelta\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"endorseVault\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"endorseVault\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\"},{\"name\":\"releaseDelta\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setBanksy\",\"inputs\":[{\"name\":\"tagger\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setBanksy\",\"inputs\":[{\"name\":\"tagger\",\"type\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"tagVault\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\"},{\"name\":\"tag\",\"type\":\"string\"}],\"outputs\":[],\"gas\":186064},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"numReleases\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":1388},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"releases\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":1533},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"numVaults\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":1663},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"vaults\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"},{\"name\":\"arg1\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":1808},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"tokens\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":1623},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"numTokens\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":1538},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"isRegistered\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":1783},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"governance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":1598},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pendingGovernance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":1628},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"tags\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"gas\":10229},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"banksy\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":1903}]"
func (s *Service) Run(cfg config.Config, ctx context.Context) {

	s.log.Info("Running token service")

	contractAbi, err := abi.JSON(strings.NewReader(myABI))
	if err != nil {
		log.Error("failed to parse contract ABI")
		return
	}

	var Contract = bind.NewBoundContract(
		common.HexToAddress(s.contractAddress.Address),
		contractAbi,
		s.eth,
		s.eth,
		s.eth,
	)

	myAddress, err := GetAddresses(Contract, s.log, common.HexToAddress(s.contractAddress.Address))
	if err != nil {
		s.log.Error("failed to call get_registry")
		return
	}

	contractAbi, err = abi.JSON(strings.NewReader(ABI))
	if err != nil {
		s.log.Error("failed to parse contract ABI")
		return
	}

	s.log.Info("myaddress=", myAddress)

	var poolContract = bind.NewBoundContract(
		common.HexToAddress(myAddress),
		contractAbi,
		s.eth,
		s.eth,
		s.eth,
	)

	contractAbi, err = abi.JSON(strings.NewReader(ABIV))
	if err != nil {
		s.log.Error("failed to parse contract ABI")
		return
	}

	var vaultContract = bind.NewBoundContract(
		common.HexToAddress(s.contractAddress.Addressv),
		contractAbi,
		s.eth,
		s.eth,
		s.eth,
	)

	count, err := GetCount(poolContract, s.log)
	if err != nil {
		s.log.Error("failed to get pool_count")
		return
	}

	intCount, err := strconv.Atoi(count)
	if err != nil {
		s.log.Error("failed to convert string to int")
		return
	}
    s.log.Info("count=",intCount)
	addresses := GetBody(s.log)

	s.log.Info("len==",len(addresses))
	for k := 0; k < len(addresses); k++ {
		address:=addresses[k]
		s.log.Info("CURRENT ADDRESS=", address)

		latestVault, err := GetLatestVault(vaultContract, s.log, address)
		if err != nil {
			s.log.Error("failed to get latest vault")
			break
		}
		s.log.Info("latest vault=", latestVault)




		token, err := s.tokenQ.FilterByAddress(address,address).Get()

		if err != nil {
			s.log.WithError(err).Error("failed to get token")
			return
		}

		if token == nil{
			s.log.WithError(err).Error(" getting zero token")
		}


		if (latestVault != "0x0000000000000000000000000000000000000000")&&(token!=nil)&&(token.Vault!=latestVault){
			err = s.tokenQ.DeleteByAddresses(token.Asset,token.Addresslp)
			if err != nil {
				s.log.Error(err, "failed to delete debtor")
				return
			}
			s.log.Info("TOKEN DELETED:",token.Asset,"  ",token.Addresslp)

			_, err = s.tokenQ.Insert(data.Token{
				Asset:      address,
				Addresslp:   address,
				Vault: latestVault,
			})
			if err != nil {
				s.log.WithError(err).Error("failed to insert data")
				return
			}
			s.log.Info("TOKEN INSERTED ",address,"  ",address,"  ",latestVault)
		}

		if (latestVault != "0x0000000000000000000000000000000000000000")&&(token==nil){
			_, err = s.tokenQ.Insert(data.Token{
				Asset:      address,
				Addresslp:   address,
				Vault: latestVault,
			})
			if err != nil {
				s.log.WithError(err).Error("failed to insert data")
				return
			}
			s.log.Info("TOKEN INSERTED ",address,"  ",address,"  ",latestVault)
		}

		//latest vault
		for i := 0; i < (intCount); i++ {

			poolAddress, err := GetPool(poolContract, s.log, i)
			if err != nil {
				s.log.Error("failed to get pool_count")
				return
			}
			s.log.Info(poolAddress)

			poolCoinsNumber, err := GetCoinsNumber(poolContract, s.log, poolAddress)
			if err != nil {
				s.log.Error("failed to get pool addresses")
				return
			}

			s.log.Info(poolCoinsNumber)
			s.log.Info(string(poolCoinsNumber[1]))

			intCoinsNumber, err := strconv.Atoi(string(poolCoinsNumber[1]))
			if err != nil {
				s.log.Error("failed to convert string to int")
				return
			}

			intCoinsUnderlineNumber, err := strconv.Atoi(string(poolCoinsNumber[3]))
			if err != nil {
				s.log.Error("failed to convert string to int")
				return
			}

			s.log.Info("intCoinsNumber=", intCoinsNumber)

			if intCoinsNumber < 4 {
				addresses, err := GetUnderlyingCoins(poolContract, s.log, poolAddress)
				if err != nil {
					s.log.Error("failed to get get_underlying_coins")
					return
				}
				s.log.Info("addresses=", addresses)

				for j := 0; j < +intCoinsUnderlineNumber; j++ {

					coinAddress := fmt.Sprintf(addresses[(j*43 + 1):((j+1)*43 + 0)])
					s.log.Info("address number ", j, "=", coinAddress)
					s.log.Info(address == coinAddress)
					if address == coinAddress {
						addressLp, err := GetLpAddress(poolContract, s.log, poolAddress)
						if err != nil {
							s.log.Error("failed to get pool_count")
							break
						}
						s.log.Info("lp_token=", addressLp)

						if addressLp != "0x0000000000000000000000000000000000000000" {
							latestVault, err = GetLatestVault(vaultContract, s.log, addressLp)
							s.log.Info("err=",err)
							if err != nil {
								s.log.Error("failed to get latest_vault")
								break
							}
							s.log.Info("latest vault=", latestVault)
						}


						token, err := s.tokenQ.FilterByAddress(address,addressLp).Get()

						if err != nil {
							s.log.WithError(err).Error("failed to get token")
							return
						}

						if token == nil{
							s.log.WithError(err).Error("getting zero token")
						}


						if (token!=nil)&&(token.Vault!=latestVault){
							err = s.tokenQ.DeleteByAddresses(address,addressLp)
							if err != nil {
								s.log.Error(err, "failed to delete debtor")
								return
							}
							s.log.Info("TOKEN DELETED ",address,"  ",addressLp)

							_, err = s.tokenQ.Insert(data.Token{
								Asset:      address,
								Addresslp:   addressLp,
								Vault: latestVault,
							})
							if err != nil {
								s.log.WithError(err).Error("failed to insert data")
								return
							}
							s.log.Info("TOKEN INSERTED ",address,"  ",addressLp,"  ",latestVault)
						}

						if token==nil{
							_, err = s.tokenQ.Insert(data.Token{
								Asset:      address,
								Addresslp:   addressLp,
								Vault: latestVault,
							})
							if err != nil {
								s.log.WithError(err).Error("failed to insert data")
								return
							}
							s.log.Info("TOKEN INSERTED ",address,"  ",addressLp,"  ",latestVault)
						}
					}
				}
			}
		}
	}
}
